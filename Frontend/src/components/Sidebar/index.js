import React from 'react';
import {
  SidebarContainer,
  Icon,
  CloseIcon,
  SidebarMenu,
  SidebarLink,
  SidebarRoute,
  SideBtnWrap
} from './SidebarElements';

const Sidebar = ({ isOpen, toggle }) => {

    function h(){
        console.log("Helloooooo")
    }
    function handleLogout(){
        localStorage.removeItem('token')
    }


  return (
    <SidebarContainer isOpen={isOpen} onClick={toggle}>
      <Icon onClick={toggle}>
        <CloseIcon />
      </Icon>
      <SidebarMenu>

        <SidebarLink onClick={h} to='/profile'>Profile</SidebarLink>
        <SidebarLink to='/'>Desserts</SidebarLink>
        <SidebarLink onClick={handleLogout} to='/'>Logout</SidebarLink>
      </SidebarMenu>
      {/*<SideBtnWrap>*/}
      {/*  <SidebarRoute to='/'>Order Now</SidebarRoute>*/}
      {/*</SideBtnWrap>*/}
    </SidebarContainer>
  );
};

export default Sidebar;
