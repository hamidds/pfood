import axios from "axios";

export default function validateInfo(values) {
    let errors = {}

    // phone number
    if (!values.phone_number.trim()) {
        errors.phone_number = "Phone number required";
    } else {
        // axios.get(`http://localhost:8000/signup/check/${values.phone_number}`).then((response) => {
        // }).catch(function (error) {
        //     // console.log(error.response.data.errors.body)
        //     // errors.phone_number = (error.response.data.errors.body).toString();
        //     errors.phone_number = "Phonenumber already exist";
        // });
        // console.log(errors)
        // const err = checkUniqueness(values.phone_number, errors)
    }

    // password
    if (!values.password) {
        errors.password = "Password is required";
    } else if (values.password.length < 8) {
        errors.password = "Password needs to be 8 character or more";
    }

    return errors;
}

async function checkUniqueness(phone_number, errors) {
    return await axios.get(`http://localhost:8000/signup/check/${phone_number}`).then((response) => {
    }).catch(function (error) {
        errors.phone_number = error.response.data.errors.body
    })
}
