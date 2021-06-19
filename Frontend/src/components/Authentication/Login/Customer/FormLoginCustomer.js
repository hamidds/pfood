import React from 'react'
import useFormCustomerLogin from './useFormCustomerLogin'
import validate from './validateCustomerLoginInfo'
import {Link} from "react-router-dom";

const FormLoginCustomer = ({submitForm}) => {
    const {handleChange, values, handleSubmit, errors}
        = useFormCustomerLogin(
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
                        Phone number:
                    </label>
                    <input
                        id="phone_number"
                        type="text"
                        name="phone_number"
                        className="form-input"
                        placeholder="Enter your phone number"
                        value={values.phone_number}
                        onChange={handleChange}
                    />
                    {errors.phone_number && <p>{errors.phone_number}</p>}
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
            Don't have an account? <Link to={"/"} exact>Register</Link>

            </span>
                <span className="form-input-manager-register">
            Login as a manager. Login <Link to={"/managerLogin"} exact>here</Link>

            </span>
            </form>
        </div>
    )
}

export default FormLoginCustomer
