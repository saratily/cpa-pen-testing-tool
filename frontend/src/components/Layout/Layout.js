import { Fragment } from 'react';

import './Layout.css';

import NavigationBar from './NavigationBar';
import MySideNav from '../MySideNav';

const Layout = (props) => {
  return (
    <Fragment>
      <NavigationBar/>
      <MySideNav />
      <main>
        <div className="container">{props.children}</div>
      </main>
    </Fragment>
  );
};

export default Layout;
