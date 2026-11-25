import { RouterProvider } from 'react-router';
import './index.css';
import { routers } from './routes/routes';

function App() {
  return <RouterProvider router={routers} />;
}

export default App;
