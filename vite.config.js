import vue from '@vitejs/plugin-vue';
import { defineConfig } from 'vite';
import svgLoader from 'vite-svg-loader';

// https://vitejs.dev/config/
export default defineConfig({
	base: '/',
	assetsInclude: ['**/*.JPG', '**/*.PNG', '**/*.PDF'],
	plugins: [vue(), svgLoader()]
});
