import React from 'react'
import useFormManagerLogin from './useFormManagerLogin'
import validate from './validateManagerLoginInfo'
import {BrowserRouter as Router, Link} from "react-router-dom";

const FormLoginManager = ({submitForm}) => {
    const {handleChange, values, handleSubmit, errors}
        = useFormManagerLogin(
        submitForm,
        validate
    );


    return (
        <div className="form-content-right">
            <div className="">

            </div>
            <form className="form" onSubmit={handleSubmit} noValidate>
                <h1>Login Page </h1>
                <div className="form-inputs">
                    <label htmlFor="phone_number" className="form-label">
                        Email address:
                    </label>
                    <input
                        id="email"
                        type="text"
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
                    Log in
                </button>
                <span className="form-input-login">
            Don't have an account? <Link to={"/managerSignup"} exact>Register</Link>
            </span>
                <span className="form-input-manager-register">
            Login as a customer. Login <Link to={"/customerLogin"} exact>here</Link>

            </span>
            </form>
        </div>
    )
}

export default FormLoginManager
