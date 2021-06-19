import React from 'react'
import useFormCustomerRegister from './useFormCustomerRegister'
import validate from './validateCustomerRegistrationInfo'
import {Link} from "react-router-dom";

const FormSignupCustomer = ({submitForm}) => {
    const {handleChange, values, handleSubmit, errors}
        = useFormCustomerRegister(
        submitForm,
        validate
    );


    return (
        <div className="form-content-right">
            <div className="">

            </div>
            <form className="form" onSubmit={handleSubmit} noValidate>
                <h1>Register Page </h1>
                <div className="form-inputs">
                    <label htmlFor="phone_number" className="form-label">
                        * Phone number:
                    </label>
                    <input
                        id="phone_number"
                        type="text"
                        name="phone_number"
                        className="form-input"
                        value={values.phone_number}
                        onChange={handleChange}
                        placeholder="Enter your phone number"
                    />
                    {errors.phone_number && <p>{errors.phone_number}</p>}
                </div>
                <div className="form-inputs">
                    <label htmlFor="password" className="form-label">
                        * Password:
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
                    Already have an account? <Link to={"/customerLogin"} exact>Login</Link>

            </span>
                <span className="form-input-manager-register">
                {/*<Router>*/}
                    Register as a Manager? Click <Link to={"/managerSignup"} exact>here</Link>
                    {/*</Router>*/}
            </span>
            </form>
        </div>
    )
}

export default FormSignupCustomer
