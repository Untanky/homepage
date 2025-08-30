import esbuild from 'esbuild';
import tailwindPlugin from 'esbuild-plugin-tailwindcss';
import { writeFile } from 'node:fs/promises'

const result = await esbuild.build({
  entryPoints: ['assets/index.ts', 'assets/index.css'],
  entryNames: '[ext]/[name]-[hash]',
  bundle: true,
  outbase: 'assets',
  outdir: 'dist',
  metafile: true,
  plugins: [
    tailwindPlugin({}),
  ],
});

const manifest = Object.entries(result.metafile.outputs)
  .filter(([_, meta]) => meta.entryPoint)
  .map(([outfile, { entryPoint }]) => ({
    entryPoint,
    outfile,
  }));

writeFile("dist/manifest.json", JSON.stringify(manifest));

