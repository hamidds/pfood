import React, {useState} from 'react';
import FormSignupManager from './FormSignupManager';
import '../../RegisterForm.css';

const FormManagerRegister = () => {
    const [isSubmited, setIsSubmitted] = useState(false);

    function submitForm() {
        setIsSubmitted(true);
    }
    return (
        <>
            <div className='form-container'>
                <span className='close-btn'>x</span>
                <div className='form-content-left'>
                    <img className='form-img' src='img/img-manager.svg' alt='manager'/>
                </div>
                {!isSubmited ? (<FormSignupManager submitForm={submitForm} />) : (console.log("form submitted!"))}
            </div>
        </>
    )
}

export default FormManagerRegister
