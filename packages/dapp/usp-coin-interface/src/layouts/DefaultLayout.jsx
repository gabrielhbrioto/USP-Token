import { Outlet, Link } from 'react-router-dom';

export function DefaultLayout() {
  return (
    <div>
      <nav style={{ padding: '1rem', background: '#eee', marginBottom: '1rem' }}>
        <Link to="/" style={{ marginRight: '1rem' }}>Início</Link>
        <Link to="/certificates">Certificados</Link>
        <Link to="/governance">Governança</Link>
        <Link to="/identity-management">Gerenciamento de Identidade</Link>
        <Link to="/paymaster-management">Gerenciamento de Paymaster</Link>
        <Link to="/system-parameters">Parâmetros do Sistema</Link>
        <Link to="/wallet">Carteira</Link>
      </nav>

      <main style={{ padding: '1rem' }}>
        <Outlet />
      </main>
    </div>
  );
}