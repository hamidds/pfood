export default function validateInfo(values) {
    let errors = {}

    // phone number
    if (!values.phone_number.trim()) {
        errors.phone_number = "Phone number required";
    }

    // password
    const passwordRegex = /^(?=.*\d)(?=.*[a-z])(?=.*[a-zA-Z]).{8,}$/;

    if (!values.password) {
        errors.password = "Password is required";
    } else if (!values.password.match(passwordRegex)) {
        errors.password = "Password needs to be at least 8 character and contains at least one numeric digit";
    }

    return errors;
}
