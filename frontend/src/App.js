import logo from './logo.svg';
import './App.css';

import { Routes, Route, Navigate } from 'react-router-dom';

import Layout from './components/Layout/Layout';
import HomePage from './pages/HomePage';
import AuthPage from './pages/AuthPage';
import PostsPage from './pages/PostsPage';
import { useContext } from 'react';
import AuthContext from './store/auth-context';

function App() {
/*
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
  */

          /*
        <Route path="/posts">
          {authContext.loggedIn && <PostsPage />}
          {!authContext.loggedIn && <Navigate to="/auth" />}
        </Route>
        <Route path="*">
          <Navigate to="/"/>
        </Route>


        <Route path="/posts">
          {authContext.loggedIn && navigate("/posts")}
          {!authContext.loggedIn && navigate("/auth")}
        </Route>
        <Route path='*' element={<Navigate to='/' />} />
                
        <Route path="*" element={<HomePage />} />


        */

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

        <Route path='*' element={<Navigate to='/' />} />
      </Routes>
    </Layout>
  );

}

export default App;
