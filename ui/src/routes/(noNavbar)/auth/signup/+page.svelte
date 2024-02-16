<script lang="ts">
    import {
        UserRegisterRequest,
        validatePassword,
        validateUsername,
        validateEmail
    } from "@encedeus/registry-js-api";
    import Input from "$lib/components/generic/Input.svelte";
    import Checkbox from "$lib/components/Checkbox.svelte";
    import Button from "$lib/components/generic/Button.svelte";
    import ErrorTextBox from "$lib/components/ErrorTextBox.svelte";
    import {goto} from "$app/navigation";
    import {accessTokenStore} from "$lib/stores/accessTokenStore";
    import {getApi} from "$lib/api/api";

    let username: string = "", email: string = "", password: string = "";
    let nameError: string = "", emailError: string = "", passwordError: string = "", responseError = "";
    let showPassword: boolean;

    const api = getApi()

    function resetErrors() {
        nameError = "";
        emailError = "";
        passwordError = "";
        responseError = "";
    }

    async function handleSignUp() {

        resetErrors();

        const nameErr = validateUsername(username);
        const emailErr = validateEmail(email);
        const passErr = validatePassword(password);

        if (nameErr) {
            nameError = nameErr.message;
        }
        if (emailErr) {
            emailError = emailErr.message;
        }
        if (passErr) {
            passwordError = passErr.message;
        }

        if (nameErr || emailErr || passErr) {
            return;
        }

        const request: UserRegisterRequest = {
            name: username,
            email: email,
            password: password
        };

        const response = await api.AuthService.SignUp(request);

        if (response.error) {
            responseError = response.error.message;
            return;
        }

        goto("/");
    }

    export function handleShowHidePassword() {
        showPassword = !showPassword;
    }

</script>

<div id="page">

    <Input type="text"
           label="name"
           placeholder="name"
           error={!!nameError}
           bind:value={username}
           bind:helperText={nameError}/>

    <Input type="text"
           label="email"
           placeholder="email"
           error={!!emailError}
           bind:value={email}
           bind:helperText={emailError}/>

    <Input type={showPassword ? "text" : "password"}
           label="password"
           placeholder="password"
           error={!!passwordError}
           bind:value={password}
           bind:helperText={passwordError}/>

    <Checkbox className="m-auto" onclick={handleShowHidePassword}>Show password</Checkbox>

    <ErrorTextBox className="m-auto" bind:value={responseError}/>

    <Button className="m-auto" onClick={handleSignUp}>Sign Up</Button>
</div>

<style>
    #page {
        text-align: center;
    }
</style>
