import logo from './logo.svg';
import './App.css';

import { useContext } from 'react';
import AuthContext from './store/auth-context';

import { Routes, Route, Navigate } from 'react-router-dom';

import Layout from './components/Layout/Layout';
import HomePage from './pages/HomePage';
import AuthPage from './pages/AuthPage';

import PostsPage from './pages/PostsPage';
import PenetrationsPage from './pages/PenetrationsPage';

import Reconnaissance from './pages/ReconnaissancePage';
import InitialExploration from './pages/InitialExplorationPage';
import NetworkScanning from './pages/NetworkScanning';
import Shodan from './pages/ShodanPage';
import ReconNg from './pages/ReconNg';

import Scanning from './pages/ScanningPage';
import WebAppEnumeration from './pages/WebAppEnumeration';
import DNSEnumeration from './pages/DNSEnumeration';


import Exploitation from './pages/ExploitationPage';
import PostExploitation from './pages/PostExploitationPage';
import Reporting from './pages/ReportingPage';


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

        <Route path="/reconnaissance/:uuid" element={authContext.loggedIn ? <Reconnaissance /> : <AuthPage /> } />
        <Route path="/initial-exploration/:uuid" element={authContext.loggedIn ? <InitialExploration /> : <AuthPage /> } />
        <Route path="/dns-enumeration/:uuid" element={authContext.loggedIn ? <DNSEnumeration /> : <AuthPage /> } />
        <Route path="/shodan/:uuid" element={authContext.loggedIn ? <Shodan /> : <AuthPage /> } />
        <Route path="/recon-ng/:uuid" element={authContext.loggedIn ? <ReconNg /> : <AuthPage /> } />

        <Route path="/scanning/:uuid" element={authContext.loggedIn ? <Scanning /> : <AuthPage /> } />
        <Route path="/network-scanning/:uuid" element={authContext.loggedIn ? <NetworkScanning /> : <AuthPage /> } />
        <Route path="/web-app-enumeration/:uuid" element={authContext.loggedIn ? <WebAppEnumeration /> : <AuthPage /> } />

        <Route path="/exploitation/:uuid" element={authContext.loggedIn ? <Exploitation /> : <AuthPage /> } />
        <Route path="/post-exploitation/:uuid" element={authContext.loggedIn ? <PostExploitation /> : <AuthPage /> } />
        <Route path="/reporting/:uuid" element={authContext.loggedIn ? <Reporting /> : <AuthPage /> } />



        <Route path='*' element={<Navigate to='/' />} />
      </Routes>
    </Layout>
  );

}

export default App;
