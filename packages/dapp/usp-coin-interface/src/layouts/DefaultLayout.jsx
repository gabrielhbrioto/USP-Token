import { Outlet, NavLink } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import './DefaultLayout.css';

export function DefaultLayout({ showNavbar }) {
  const { logout } = useAuth();

  return (
    <div className="usp-layout-root">
      {showNavbar && (
        <nav className="usp-navbar">
          <div className="usp-navbar-inner">
            <div className="usp-nav-links-container">
              <button type="button" className="usp-nav-link usp-nav-button" onClick={logout}>
                Logout
              </button>
              <NavLink to="/certificates" className={({ isActive }) => `usp-nav-link${isActive ? ' usp-nav-link-active' : ''}`}>
                Certificados
              </NavLink>
              <NavLink to="/wallet" className={({ isActive }) => `usp-nav-link${isActive ? ' usp-nav-link-active' : ''}`}>
                Carteira
              </NavLink>
            </div>
          </div>
        </nav>
      )}

      <main className={`usp-main${showNavbar ? ' usp-main-with-navbar' : ''}`}>
        <Outlet />
      </main>
    </div>
  );
}