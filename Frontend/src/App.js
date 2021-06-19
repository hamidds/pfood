import './App.css';
import {BrowserRouter as Router, Route, Switch} from "react-router-dom";
import FormCustomerRegister from './components/FormCustomerRegister';
import React from 'react';
import {GlobalStyle} from './globalStyles';
import Hero from './components/Hero';
import Products from './components/Products';
import Footer from './components/Footer';
import axios from "axios";

function App() {
    function getFoods() {
        const afoods = axios.get(`http://localhost:8000/foodss`).then((response) => {
            console.log(response.data.foods)
            console.log("===========================================")
            return response.data.foods
        }).catch(function (error) {
        });
        console.log(afoods instanceof Object)
        let result = Object.keys(afoods).map((key) => [String(key), afoods[key]]);
        return result;
    }

    return (
        <Router>
            <Switch>
                <Route exact path="/">
                    <div className="App">
                        <FormCustomerRegister/>
                    </div>
                </Route>
                <Route exact path="/main">
                    <GlobalStyle/>
                    <Hero/>
                    <Products heading='Find your favorite' data={getFoods()}/>
                    <Footer/>
                </Route>
            </Switch>
        </Router>

    );
}

export default App;

