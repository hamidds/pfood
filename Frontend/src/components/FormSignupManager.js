import React from 'react'
import useForm from './useForm'
import validate from './validateRegistrationInfo'

const FormSignupManager = ({ submitForm }) => {
    const {handleChange, values, handleSubmit, errors}
        = useForm(
        submitForm,
        validate
    );


    return (
        <div className="form-content-right">
            <form className="form" onSubmit={handleSubmit} noValidate>
                <h1>Register Page </h1>
                <div className="form-inputs">
                    <label htmlFor="phone_number" className="form-label">
                        Phone number:
                    </label>
                    <input
                        id="email"
                        type="email"
                        name="email"
                        className="form-input"
                        placeholder="username@domain.com"
                        value={values.email}
                        onChange={handleChange}
                    />
                    {errors.email && <p>{errors.email}</p>}
                </div>
                <div className="form-inputs">
                    <label htmlFor="email" className="form-label">
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
            Already have an account? <a href="#">Login</a>
            </span>
                <span className="form-input-manager-register">
            Register as a company owner <a href="#">Here</a>
            </span>
            </form>
        </div>
    )
}

export default FormSignupManager
