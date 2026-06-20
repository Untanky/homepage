import esbuild from "esbuild";
import tailwindPlugin from "esbuild-plugin-tailwindcss";

await esbuild.build({
  entryPoints: [
    "web/main.css",
  ],
  outdir: "tmp/web",
  bundle: true,
  plugins: [
    tailwindPlugin({
      /* options */
    }),
  ],
});
