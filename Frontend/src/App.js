import './App.css';
import {BrowserRouter as Router, Route, Switch} from "react-router-dom";
import FormCustomerRegister from './components/Authentication/Signup/Customer/FormCustomerRegister';
import FormManagerRegister from "./components/Authentication/Signup/Manager/FormManagerRegister";
import FormManagerLogin from "./components/Authentication/Login/Manager/FormManagerLogin";
import FormCustomerLogin from "./components/Authentication/Login/Customer/FormCustomerLogin";
import ProfileForm from "./components/Profile/ProfileForm";
import React, {useState} from 'react';
import {GlobalStyle} from './globalStyles';
import Hero from './components/Hero';
import Products from './components/Products';
import Footer from './components/Footer';
import axios from "axios";
// import {productData} from "./components/Products/data";

function App() {
    const [user, setUser] = useState();

    async function getFoods() {
        const foods = await axios.get(`http://localhost:8000/foodss`).then((response) => {
            console.log(response.data.foods)
            console.log("===========================================")
            return response.data.foods
        }).catch(function (error) {
        });
        console.log(foods)
        return foods;
    }

    const callbackFunction = (childData) => {
        setUser(childData)
        console.log(childData)
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
                    <Products heading='Find your favorite'
                              data={[]}/>
                    <Footer/>
                </Route>
                <Route exact path="/managerSignup">
                    <div className="App">
                        <FormManagerRegister/>
                    </div>
                </Route>
                <Route exact path="/managerLogin">
                    <div className="App">
                        <FormManagerLogin/>
                    </div>
                </Route>
                <Route exact path="/customerLogin">
                    <div className="App">
                        <FormCustomerLogin/>
                    </div>
                </Route>
                <Route exact path="/profile">
                    <div className="App">
                        <ProfileForm/>
                    </div>
                </Route>
            </Switch>
        </Router>

    );
}

export default App;

