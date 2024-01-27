import { defineConfig, loadEnv } from 'vite';
import react from '@vitejs/plugin-react';

// https://vitejs.dev/config/
export default ({ mode }) => {
  process.env = { ...process.env, ...loadEnv(mode, process.cwd()) };

  return defineConfig({
    plugins: [react()],
    server: {
      host: '0.0.0.0',
      hmr: {
        clientPort: parseInt(process.env.VITE_WEB_PORT) || 5173,
      },
      port: 5173,
      watch: {
        usePolling: true,
      },
    },
  });
};
