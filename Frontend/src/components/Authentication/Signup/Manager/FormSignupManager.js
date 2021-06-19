import React from 'react'
import useFormManagerRegister from '../Manager/useFormManagerRegister'
import validate from '../Manager/validateManagerRegisterationInfo'
import {BrowserRouter as Router, Link, Redirect} from "react-router-dom";

const FormSignupManager = ({ submitForm }) => {
    const {handleChange, values, handleSubmit, errors}
        = useFormManagerRegister(
        submitForm,
        validate
    );


    return (
        <div className="form-content-right">
            <form className="form" onSubmit={handleSubmit} noValidate>
                <h1>Register Page </h1>
                <div className="form-inputs">
                    <label htmlFor="email" className="form-label">
                        Email address:
                    </label>
                    <input
                        id="email"
                        type="email"
                        name="email"
                        className="form-input"
                        placeholder="Enter your email address"
                        value={values.email}
                        onChange={handleChange}
                    />
                    {errors.email && <p>{errors.email}</p>}
                </div>
                <div className="form-inputs">
                    <label htmlFor="password" className="form-label">
                        Password:
                    </label>
                    <input
                        id="password"
                        type="password"
                        name="password"
                        className="form-input"
                        placeholder="Enter your password"
                        value={values.password}
                        onChange={handleChange}
                    />
                    {errors.password && <p>{errors.password}</p>}
                </div>
                <button className="form-input-btn" type="submit">
                    Sign up
                </button>
                <span className="form-input-login">
                    Already have an account? <Link to={"/managerLogin"} exact>Login</Link>
            </span>
                <span className="form-input-manager-register">
                {/*<Router>*/}
                    Register as a customer? Click <Link to={"/"} exact>here</Link>
                {/*</Router>*/}
            </span>
            </form>
        </div>
    )
}

export default FormSignupManager
