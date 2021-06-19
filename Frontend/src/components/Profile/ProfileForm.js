import React, {useEffect, useState} from 'react';
import FormProfile from './FormProfile';
import './ProfileForm.css';
import {Link, Redirect} from "react-router-dom";
import axios from "axios";

const api_url = `http://localhost:8000`

const ProfileForm = () => {
    const [isSubmitted, setIsSubmitted] = useState(false);
    const [name, setName] = useState('Customer Name');

    useEffect(() => {
        axios.get(`${api_url}/profile`).then((response) => {
            localStorage.setItem("token", response.data.user.token)
            console.log(response.data.user.token)
            setName(response.data.user.name)
        }).catch(function (error) {
            // handle error
            console.log("profile name error")
        });

        // setName("Hamidreza")
    }, [])

    function submitForm() {
        setIsSubmitted(true);
    }

    return (
        <>
            <div className='form-container'>
                <span className='close-btn'>
                    <Link to={"/main"} exact>x</Link>
                </span>
                <div className='form-content-left'>
                    <p>Hi, {name}</p>
                    <img className='form-img' src='img/img-customer.svg' alt='prfoile-image'/>
                </div>
                {!isSubmitted ? (<FormProfile submitForm={submitForm}/>) : <Redirect to="/main"/>}
            </div>
        </>
    )
}

export default ProfileForm

