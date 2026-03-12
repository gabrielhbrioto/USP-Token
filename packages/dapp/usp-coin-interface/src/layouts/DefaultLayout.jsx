import { Outlet, Link } from 'react-router-dom';

export function DefaultLayout({ showNavbar }) {
  return (
    <div>
      {/* A barra de navegação só renderiza se showNavbar for true */}
      {showNavbar && (
        <nav style={{ padding: '1rem', background: '#eee', marginBottom: '1rem' }}>
          <Link to="/" style={{ marginRight: '1rem' }}>Início</Link>
          <Link to="/certificates" style={{ marginRight: '1rem' }}>Certificados</Link>
          <Link to="/governance" style={{ marginRight: '1rem' }}>Governança</Link>
          <Link to="/identity-management" style={{ marginRight: '1rem' }}>Gerenciamento de Identidade</Link>
          <Link to="/paymaster-management" style={{ marginRight: '1rem' }}>Gerenciamento de Paymaster</Link>
          <Link to="/system-parameters" style={{ marginRight: '1rem' }}>Parâmetros do Sistema</Link>
          <Link to="/wallet">Carteira</Link>
        </nav>
      )}

      <main style={{ padding: '1rem' }}>
        <Outlet />
      </main>
    </div>
  );
}