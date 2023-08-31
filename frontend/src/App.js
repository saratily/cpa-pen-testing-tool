import logo from './logo.svg';
import './App.css';

import { useContext } from 'react';
import AuthContext from './store/auth-context';

import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';

import Layout from './components/Layout/Layout';
import HomePage from './pages/HomePage';
import AuthPage from './pages/AuthPage';

import PostsPage from './pages/PostsPage';
import PenetrationsPage from './pages/PenetrationsPage';



import Planning from './pages/PlanningPage';
import Scanning from './pages/ScanningPage';

function App() {

  const authContext = useContext(AuthContext);

  return (
    <Layout>    
      <Routes>
        <Route path="/" exact element={<HomePage/>} />
        {!authContext.loggedIn && (
          <Route path="/auth" element={<AuthPage />} />
        )}
        <Route path="/posts" element=
          {authContext.loggedIn ? <PostsPage /> : <AuthPage /> } />
        <Route path="/penetrations" element=
          {authContext.loggedIn ? <PenetrationsPage /> : <AuthPage /> } />

        <Route path="/planning" element={<Planning />} /> 
        <Route path="/scanning" element={<Scanning />} /> 


        <Route path='*' element={<Navigate to='/' />} />
      </Routes>
    </Layout>
  );

}

export default App;
