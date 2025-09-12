import React from 'react';
import './Header.css';

const Header = () => {
    return (
        <header className="header">
            <h1>Inventory Management System</h1>
            <nav>
                <ul>
                    <li><a href="/">Houses</a></li>
                    <li><a href="/rooms">Rooms</a></li>
                </ul>
            </nav>
        </header>
    );
};

export default Header;
