import { useState } from 'react';
import { BrowserRouter as Router, Route, Switch, Link } from 'react-router-dom';
import './App.css';
import Home from './components/Home';  // Assume this is your Home component
import Library from './components/Library';  // Your Library view
import BookDetails from './components/BookDetails';  // Book details page

function App() {
    return (
        <Router>
            <nav>
                <Link to="/">Home</Link>
                <Link to="/library">Library</Link>
                {/* Add more navigation links as needed */}
            </nav>

            <Switch>
                <Route exact path="/" component={Home} />
                <Route path="/library" component={Library} />
                <Route path="/book/:id" component={BookDetails} />
                {/* Define other routes here */}
            </Switch>
        </Router>
    );
}

export default App;
