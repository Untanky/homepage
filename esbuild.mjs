import esbuild from 'esbuild';
import tailwindPlugin from 'esbuild-plugin-tailwindcss';
import { lookup } from 'mime-types';
import { createHash, Hash } from 'node:crypto';
import { createReadStream, createWriteStream } from 'node:fs';
import { stat, writeFile } from 'node:fs/promises';
import { PassThrough } from 'node:stream';
import { pipeline } from 'node:stream/promises';
import { constants, createGzip, createBrotliCompress, createDeflate } from 'node:zlib';

/**
 * Remove the prefix of a string
 *
 * @param {string} str the string to trim
 * @param {string} prefix the prefix to remove
 *
 * @return {string} if `prefix` is prefix of `str`, `prefix` is removed from `str` otherwise `str` is returned as is
 */
function trimPrefix(str, prefix) {
  if (str.startsWith(prefix)) {
    return str.slice(prefix.length);
  }
  return str;
}

/**
 * @typedef EntryVersion {{ filepath: string; encoding: 'none' | 'gzip' | 'br' | 'deflate', size: number }}
 * @typedef Entry {{ entrypoint: string; etag: string; mimetype: string; versions: EntryVersion[] }}
 */

/**
 * Transform output files
 *
 * Output files are precompressed to serve optimally at runtime and an etag
 * is calculated to optimize caching.
 *
 * @param outfile {string} the path to the outputfile
 * @param entrypoint {string} the path to the entrypoint 
 * 
 * @return {Promise<Entry>}
 */
async function transformOutfile(outfile, entrypoint) {
  const hash = createHash('SHA3-256');
  await pipeline(createReadStream(outfile), hash);
  const etag = hash.digest('base64');

  const filenameForEncoding = {
    none: outfile,
    brotli: `${outfile}.br`,
    gzip: `${outfile}.gz`,
    deflate: `${outfile}.zz`,
  };

  const transformForEncoding = {
    brotli: createBrotliCompress({
      params: {
        [constants.BROTLI_PARAM_QUALITY]: 11,
      },
    }),
    gzip: createGzip({ level: 9 }),
    deflate: createDeflate({ level: 9 }),
  };

  const fileReadStream = createReadStream(outfile);

  const versions = await Promise.all(Object.keys(filenameForEncoding)
    .map((encoding) => {
      const passthrough = new PassThrough();
      fileReadStream.pipe(passthrough);
      return [encoding, passthrough];
    })
    .map(async ([encoding, writer]) => {
      const transform = transformForEncoding[encoding];

      const filepath = filenameForEncoding[encoding];
      if (transform) {
        await pipeline(writer, transform, createWriteStream(filepath));
      }

      const { size } = await stat(filepath);
      return {
        filepath: trimPrefix(filepath, 'internal/assets/dist/'),
        encoding,
        size,
      };
    }));

  const mimetype = lookup(outfile) || 'octet-stream';

  return {
    entrypoint: trimPrefix(entrypoint, 'assets/'),
    etag,
    mimetype,
    versions,
  };
}

const result = await esbuild.build({
  entryPoints: ['assets/index.ts', 'assets/index.css'],
  entryNames: '[ext]/[name]-[hash]',
  bundle: true,
  outbase: 'assets',
  outdir: 'internal/assets/dist/assets',
  metafile: true,
  plugins: [
    tailwindPlugin({}),
  ],
});

const manifest = await Promise.all(Object.entries(result.metafile.outputs)
  .filter(([_, meta]) => meta.entryPoint)
  .map(([outfile, { entryPoint }]) => transformOutfile(outfile, entryPoint)));

await writeFile("./dist/manifest.json", JSON.stringify(manifest));

