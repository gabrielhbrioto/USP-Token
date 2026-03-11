import { createBrowserRouter } from 'react-router-dom';
import { DefaultLayout } from '../layouts/DefaultLayout';
import { Home } from '../pages/Home';
import { Authentication } from '../pages/Authentication';
import { Certificates } from '../pages/Certificates';
import { Governance } from '../pages/Governance';
import { IdentityManagement } from '../pages/IdentityManagement';
import { PaymasterManagement } from '../pages/PaymasterManagement';
import { SystemParameters } from '../pages/SystemParameters';
import { Wallet } from '../pages/Wallet';

export const router = createBrowserRouter([
  {
    path: '/',
    element: <DefaultLayout />, // O Layout é o elemento pai
    children: [
      {
        path: '/', // Quando a URL for apenas "/", renderiza a Home dentro do Layout
        element: <Home />,
      },
      {
        path: '/auth', // Quando for "/auth", renderiza o Authentication dentro do Layout
        element: <Authentication />,
      },
      {
        path: '/certificates', // Quando for "/certificates", renderiza o Certificates dentro do Layout
        element: <Certificates />,
      },
      {
        path: '/governance', // Quando for "/governance", renderiza o Governance dentro do Layout
        element: <Governance />,
      },
      {
        path: '/identity-management', // Quando for "/identity-management", renderiza o IdentityManagement dentro do Layout
        element: <IdentityManagement />,
      },
      {
        path: '/paymaster-management', // Quando for "/paymaster-management", renderiza o PaymasterManagement dentro do Layout
        element: <PaymasterManagement />,
      },
      {
        path: '/system-parameters', // Quando for "/system-parameters", renderiza o SystemParameters dentro do Layout
        element: <SystemParameters />,
      },
      {
        path: '/wallet', // Quando for "/wallet", renderiza o Wallet dentro do Layout
        element: <Wallet />,
      },
    ],
  },
]);