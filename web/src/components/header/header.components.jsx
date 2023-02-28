import React from 'react';
import './header.styles.css';

const Header = () => {
  return (
    <div className='header'>
        <div className="header-left">
          <h2>Frappy</h2>
        </div>
        <div className="header-right">
            <div className="option">
                <ul>
                    <li>Feature</li>
                    <li>Buy</li>
                    <li>Create</li>
                    <li>Marketplace</li>
                </ul>
            </div>
            <div className="connect-opt">
                <p>Connect</p>
            </div>
        </div>
    </div>
  )
}

export default Header;