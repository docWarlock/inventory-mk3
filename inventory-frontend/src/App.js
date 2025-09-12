import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Header from './components/Header';
import HouseList from './components/HouseList';
import HouseForm from './components/HouseForm';
import RoomList from './components/RoomList';
import RoomForm from './components/RoomForm';
import './App.css';

function App() {
  return (
    <Router>
      <div className="App">
        <Header />
        <main>
          <Routes>
            <Route path="/" element={<HouseList />} />
            <Route path="/houses" element={<HouseList />} />
            <Route path="/houses/new" element={<HouseForm />} />
            <Route path="/houses/:id/edit" element={<HouseForm />} />
            <Route path="/houses/:id/rooms" element={<RoomList />} />
            <Route path="/houses/:id/rooms/new" element={<RoomForm />} />
            <Route path="/rooms" element={<RoomList />} />
            <Route path="/rooms/new" element={<RoomForm />} />
            <Route path="/rooms/:id/edit" element={<RoomForm />} />
          </Routes>
        </main>
      </div>
    </Router>
  );
}

export default App;
