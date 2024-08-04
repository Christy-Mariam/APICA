import './App.css';
import CacheInput from './components/CacheInput/CacheInput';
import CacheDisplay from './components/CacheDisplay/CacheDisplay';

function App() {
  return (
    <div className="container">
      <h1>LRU Cache Management</h1>
      <CacheInput />
      <CacheDisplay />
    </div>
  );
}

export default App;
