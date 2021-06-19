import React from 'react'
import useProfileForm from './useProfileForm'
import validate from './validateCustomerRegistrationInfo'

const FormProfile = ({submitForm}) => {
    const {handleChange, values, handleSubmit, errors}
        = useProfileForm(
        submitForm,
        validate
    );

    return (
        <div className="form-content-right">
            <div className="">

            </div>
            <form className="form" onSubmit={handleSubmit} noValidate>
                <h1>Profile Page </h1>

                <div className="form-inputs">
                    <label htmlFor="name" className="form-label">
                        Name:
                    </label>
                    <input
                        id="name"
                        type="text"
                        name="name"
                        className="form-input"
                        placeholder="Enter your name"
                        value={values.name}
                        onChange={handleChange}
                    />
                </div>
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
                        // onChange={handleChange}
                        readOnly
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

                <div className="form-inputs">
                    <label htmlFor="state" className="form-label">
                        District:
                    </label>
                    <input
                        id="state"
                        type="text"
                        name="state"
                        className="form-input"
                        placeholder="Enter your district"
                        value={values.district}
                        onChange={handleChange}
                    />
                </div>

                <div className="form-inputs">
                    <label htmlFor="address" className="form-label">
                        Address:
                    </label>
                    <textarea
                        id="address"
                        type="text"
                        name="address"
                        className="form-textarea"
                        placeholder="Enter your address"
                        value={values.address}
                        onChange={handleChange}
                    />
                </div>

                <button className="form-input-btn" type="submit">
                    Save changes
                </button>

            </form>
        </div>
    )
}

export default FormProfile
