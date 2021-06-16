import './App.css';
import {BrowserRouter as Router, Route, Switch} from "react-router-dom";
import FormCustomerRegister from './components/FormCustomerRegister';

function App() {
    return (
        <Router>
            <Switch>
                <Route exact path="/">
                    <div className="App">
                        <FormCustomerRegister/>
                    </div>
                </Route>
                <Route path="/main">
                    <div className="App">
                        Main Page
                    </div>
                </Route>
            </Switch>
        </Router>

    );
}

export default App;
