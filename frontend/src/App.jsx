import { useState } from 'react';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import './App.css';
import Home from './components/Home'; 
import Library from './components/Library';
import BookDetails from './components/BookDetails';

function App() {
    return (
        <Router>
            <nav>
                <Link to="/">Home</Link>
                <Link to="/library">Library</Link>
                {/* Add more navigation links as needed */}
            </nav>

            <Routes>
                <Route exact path="/" element={<Home />} />
                <Route path="/library" element={<Library />} />
                <Route path="/book/:id" element={<BookDetails />} />
                {/* Define other routes here */}
            </Routes>
        </Router>
    );
}

export default App;

