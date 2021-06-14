import React, {useState} from 'react';
import FormSignupCustomer from './FormSignupCustomer';
import './Form.css';

const FormCustomerRegister = () => {
    const [isSubmited, setIsSubmitted] = useState(false);

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
            {!isSubmited ? (<FormSignupCustomer submitForm={submitForm} />) : console.log("submitted") }
        </div>

        </>
    )
}

export default FormCustomerRegister
