export default function validateInfo(values) {
    let errors = {}

    // email
    if (!values.email.trim()) {
        errors.email = "Email required";
    }

    // password
    const passwordRegex = /^(?=.*\d)(?=.*[a-z])(?=.*[a-zA-Z]).{8,}$/;
    if (!values.password) {
        errors.password = "Password is required";
    } else if (!values.password.match(passwordRegex)) {
        errors.password = "Password needs to be 8 character or more";
    }

    return errors;
}