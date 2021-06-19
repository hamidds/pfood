import React, {useState} from 'react';
import FormLoginManager from './FormLoginManager';
import '../../RegisterForm.css';
import {Redirect} from "react-router-dom";

const FormManagerLogin = () => {
    const [isSubmitted, setIsSubmitted] = useState(false);


    function submitForm() {
        setIsSubmitted(true);
    }
    return (
        <>
        <div className='form-container'>
            <span className='close-btn'>x</span>
            <div className='form-content-left'>
                <img className='form-img' src='img/img-Manager.svg' alt='Manager'/>
            </div>
            {!isSubmitted ? (<FormLoginManager submitForm={submitForm} />) : <Redirect to="/main"/> }
        </div>
        </>
    )
}

export default FormManagerLogin
