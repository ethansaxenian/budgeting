import { extendTheme } from '@chakra-ui/react';

export const theme = extendTheme({
  styles: {
    global: {
      '*, *::before, *::after': {
        boxSizing: 'border-box',
      },
      '*': {
        margin: 0,
      },
      'html, body': {
        height: '100%',
      },
      'img, picture, video, canvas, svg': {
        display: 'block',
        maxWidth: '100%',
      },
    },
  },
});
