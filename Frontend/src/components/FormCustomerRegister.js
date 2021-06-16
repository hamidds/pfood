import React, {useState} from 'react';
import FormSignupCustomer from './FormSignupCustomer';
import './Form.css';
import {Redirect} from "react-router-dom";

const FormCustomerRegister = () => {
    const [isSubmitted, setIsSubmitted] = useState(false);


    function submitForm() {
        setIsSubmitted(true);
    }
    return (
        <>
        <div className='form-container'>
            <span className='close-btn'>x</span>
            <div className='form-content-left'>
                <img className='form-img' src='img/img-customer.svg' alt='customer'/>
            </div>
            {!isSubmitted ? (<FormSignupCustomer submitForm={submitForm} />) : <Redirect to="/main"/> }
        </div>
        </>
    )
}

export default FormCustomerRegister
