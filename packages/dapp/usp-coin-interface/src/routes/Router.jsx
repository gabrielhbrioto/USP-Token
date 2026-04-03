import { createBrowserRouter } from 'react-router-dom';
import { DefaultLayout } from '../layouts/DefaultLayout';
import { Login } from '../pages/Login';
import { Authentication } from '../pages/Authentication';
import { Certificates } from '../pages/Certificates';
// import { Governance } from '../pages/Governance';
// import { IdentityManagement } from '../pages/IdentityManagement';
// import { PaymasterManagement } from '../pages/PaymasterManagement';
// import { SystemParameters } from '../pages/SystemParameters';
import { Wallet } from '../pages/Wallet';
// import { Dashboard } from '../pages/Dashboard';
import { ProtectedRoute } from './ProtectedRoutes';
import { AuthProvider } from '../contexts/AuthContext';
import { Outlet } from 'react-router-dom';


const RootLayout = () => {
  return (
    <AuthProvider>
      <Outlet />
    </AuthProvider>
  );
};

export const router = createBrowserRouter([
  {
    element: <RootLayout />,
    children: [
      // 1º Grupo: Sem Navbar (Home/Login e Auth)
      {
        path: '/',
        element: <DefaultLayout showNavbar={false} />, // Passando uma prop se quiser reaproveitar
        children: [
          { path: '/', element: <Login /> },
          { path: '/auth', element: <ProtectedRoute><Authentication /></ProtectedRoute> },
        ],
      },
      // 2º Grupo: Com Navbar (Todas as outras rotas protegidas)
      {
        element: <DefaultLayout showNavbar={true} />, 
        children: [
          // { 
          //   path: '/dashboard', 
          //   element: 
          //     <ProtectedRoute>
          //       <Dashboard />
          //     </ProtectedRoute> 
          // },
          { 
            path: '/certificates', 
            element: 
              <ProtectedRoute>
                <Certificates />
              </ProtectedRoute> 
          },
          // { 
          //   path: '/governance', 
          //   element: 
          //     <ProtectedRoute>
          //       <Governance />
          //     </ProtectedRoute> 
          // },
          // { 
          //   path: '/identity-management', 
          //   element: 
          //     <ProtectedRoute>
          //       <IdentityManagement />
          //     </ProtectedRoute> 
          // },
          // { 
          //   path: '/paymaster-management', 
          //   element: 
          //     <ProtectedRoute>
          //       <PaymasterManagement /></ProtectedRoute> 
          // },
          // { 
          //   path: '/system-parameters', 
          //   element: 
          //     <ProtectedRoute>
          //       <SystemParameters />
          //     </ProtectedRoute> 
          // },
          { 
            path: '/wallet', 
            element: 
              <ProtectedRoute>
                <Wallet />
              </ProtectedRoute> 
          },
        ],
      },
    ],
  },
]);