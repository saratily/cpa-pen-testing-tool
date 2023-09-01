import SideNav, {Toggle, NavItem, NavIcon, NavText} from '@trendmicro/react-sidenav';

import '@trendmicro/react-sidenav/dist/react-sidenav.css';
import { useLocation, useNavigate } from 'react-router-dom';

function MySideNav() {
    const navigate = useNavigate();
    const location = useLocation();
    const uuid = obj => obj.uuid === location.pathname.split("/")[2];
    return (
        <SideNav className='mysidenav'
            onSelect={(selected) => {
                console.log(selected);
                navigate('/'+selected+'/'+uuid);
            }}
            >
            <SideNav.Toggle />
            <SideNav.Nav defaultSelected="planning">
                <NavItem eventKey="planning">
                    <NavIcon>
                        <i className='fa-solid fa-newspaper' style={{ fontSize: "1.5em" }} />
                    </NavIcon>
                    <NavText>0. Planninng</NavText>
                </NavItem>
                <NavItem eventKey="recommaissance">
                    <NavIcon>
                        <i className='fa-solid fa-binoculars' style={{ fontSize: "1.5em" }} />
                    </NavIcon>
                    <NavText>1. Reconnaissance</NavText>
                    <NavItem eventKey="netwrok-scanning">
                        <NavText>Network Scanninng</NavText>
                    </NavItem>
                    <NavItem eventKwy="web-app-enumeration">
                        <NavText>Web App Enumeration</NavText>
                    </NavItem>
                    <NavItem eventKey="dns-enumeration">
                        <NavText>DNS Enumeration</NavText>
                    </NavItem>
                </NavItem>
                <NavItem eventKey="scanning">
                    <NavIcon>
                        <i className='fa-solid fa-magnifying-glass' style={{ fontSize: "1.5em" }} />
                    </NavIcon>
                    <NavText>2. Scanninng</NavText>
                </NavItem>
                <NavItem eventKey="expolitation">
                    <NavIcon>
                        <i className='fa-solid fa-explosion' style={{ fontSize: "1.5em" }} />
                    </NavIcon>
                    <NavText>3. Edxploitation</NavText>
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