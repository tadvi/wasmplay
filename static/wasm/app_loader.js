
/*
Important: add following lines to page.html

	<script src="wasm/wasm_exec.js"></script>
	<script src="wasm/app_loader.js"></script>
*/

if (!WebAssembly.instantiateStreaming) { // polyfill
	WebAssembly.instantiateStreaming = async (resp, importObject) => {
		const source = await (await resp).arrayBuffer();
		return await WebAssembly.instantiate(source, importObject);
	};
}

async function run() {
	console.clear();
	await go.run(inst);
	inst = await WebAssembly.instantiate(mod, go.importObject); // reset instance
}

const go = new Go();
let mod, inst;
WebAssembly.instantiateStreaming(fetch("wasm/main.wasm"), go.importObject).then((result) => {
	mod = result.module;
	inst = result.instance;

	run();
}).catch((err) => {
	console.error(err);
});
