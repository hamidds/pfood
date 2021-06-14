import React, {useState} from 'react';
import FormSignup from './FormSignup';
import './Form.css';

const Form = () => {
    const [isSubmited, setIsSubmitted] = useState(false);

    function submitForm() {
        setIsSubmitted(true);
    }
    return (
        <>
        <div className='form-container'>
            <span className='close-btn'>x</span>
            <div className='form-content-left'>
                <img className='form-img' src='img/img-2.svg' alt='spaceship'/>
            </div>
            {!isSubmited ? (<FormSignup submitForm={submitForm} />) : (console.log("form submitted!"))}
        </div>
        </>
    )
}

export default Form
