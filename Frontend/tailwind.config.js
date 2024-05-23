/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}', './node_modules/flowbite-svelte/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        'primary': '#edede9',
        'secondary': '#d5bdaf',
        'tertiary': '#e3d5ca',

      },
    }
  },
  plugins: [require('flowbite/plugin')],
};