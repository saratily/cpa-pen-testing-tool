import SideNav, {Toggle, NavItem, NavIcon, NavText} from '@trendmicro/react-sidenav';

import '@trendmicro/react-sidenav/dist/react-sidenav.css';
import { useLocation, useNavigate } from 'react-router-dom';

function MySideNav() {
    const navigate = useNavigate();
    const location = useLocation();
    const uuid = location.pathname.split("/")[2];

    return (
        <SideNav className='mysidenav'
            onSelect={(selected) => {
                console.log(selected);
                navigate('/'+selected+'/'+uuid);
            }}
            >
            <SideNav.Toggle />
            <SideNav.Nav defaultSelected="planning">

                <NavItem eventKey="recommaissance">
                    <NavIcon>
                        <i className='fa-solid fa-binoculars' style={{ fontSize: "1.5em" }} />
                    </NavIcon>
                    <NavText>Planning & Reconnaissance</NavText>
                    <NavItem eventKey="initial-exploration">
                        <NavText>Initial Exploration</NavText>
                    </NavItem>
                    <NavItem eventKey="dns-enumeration">
                        <NavText>DNS Enumeration</NavText>
                    </NavItem>

                    <NavItem eventKey="shodan">
                        <NavText>Shodan</NavText>
                    </NavItem>

                </NavItem>
                <NavItem eventKey="scanning">
                    <NavIcon>
                        <i className='fa-solid fa-magnifying-glass' style={{ fontSize: "1.5em" }} />
                    </NavIcon>
                    <NavText>Scanninng & Enumeration</NavText>
                    <NavItem eventKey="network-scanning">
                        <NavText>Network Scanninng</NavText>
                    </NavItem>
                    <NavItem eventKey="web-app-enumeration">
                        <NavText>Web App Enumeration</NavText>
                    </NavItem>
                </NavItem>
                <NavItem eventKey="exploitation">
                    <NavIcon>
                        <i className='fa-solid fa-explosion' style={{ fontSize: "1.5em" }} />
                    </NavIcon>
                    <NavText>3. Exploitation</NavText>
                </NavItem>
                <NavItem eventKey="post-exploitation">
                    <NavIcon>
                        <i className='fa-solid fa-house-crack' style={{ fontSize: "1.5em" }} />
                    </NavIcon>
                    <NavText>4. Post Exploitation</NavText>
                </NavItem>
                <NavItem eventKey="reporting">
                    <NavIcon>
                        <i className='fa-solid fa-download' style={{ fontSize: "1.5em" }} />
                    </NavIcon>
                    <NavText>%. Reporting</NavText>
                </NavItem>
            </SideNav.Nav>

        </SideNav>
        );
}

export default MySideNav;