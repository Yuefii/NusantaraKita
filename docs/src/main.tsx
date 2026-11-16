import { inject } from '@vercel/analytics';
import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { RouterProvider } from 'react-router';
import { FixLeafletIcons } from './components/ui/fix-leaflet-icon.tsx';
import { ResponseTypeProvider } from './components/ui/response-type.tsx';
import { OverviewProvider } from './context/overview-provider.tsx';
import ReactQueryProvider from './context/react-query-provider.tsx';
import { ThemeProvider } from './context/theme-provider.tsx';
import './index.css';
import { routers } from './routes/routes.ts';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ThemeProvider
      defaultTheme="system"
      storageKey="nusantara-theme"
    >
      <ReactQueryProvider>
        <ResponseTypeProvider>
          <OverviewProvider>
            <FixLeafletIcons />
            <RouterProvider router={routers} />
          </OverviewProvider>
        </ResponseTypeProvider>
      </ReactQueryProvider>
    </ThemeProvider>
  </StrictMode>,
);

inject();
