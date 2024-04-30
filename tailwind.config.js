const colors = require('tailwindcss/colors')
const defaultTheme = require('tailwindcss/defaultTheme')

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./html/**/*.go",
  ],
  theme: {
    extend: {
      colors: {
        primary: colors.cyan,
      },
      fontFamily: {
        mono: [...defaultTheme.fontFamily.mono],
        sans: [...defaultTheme.fontFamily.sans],
        serif: ['Charter', ...defaultTheme.fontFamily.serif],
      },
      typography: (theme) => ({
        DEFAULT: {
          css: {
            'code::before': {
              content: 'none',
            },
            'code::after': {
              content: 'none',
            }
          }
        }
      })
    }
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}
