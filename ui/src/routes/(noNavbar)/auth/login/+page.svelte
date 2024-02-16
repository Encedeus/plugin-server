<script lang="ts">
    import {
        UserSignInRequest,
        validatePassword,
        validateUserIdentifier
    } from "@encedeus/registry-js-api";
    import Input from "$lib/components/generic/Input.svelte";
    import Checkbox from "$lib/components/Checkbox.svelte";
    import Button from "$lib/components/generic/Button.svelte";
    import ErrorTextBox from "$lib/components/ErrorTextBox.svelte";
    import {goto} from "$app/navigation";
    import {getApi} from "$lib/api/api";

    const api = getApi();

    let password: string = "", userIdentifier: string = "";
    let passwordError: string = "", userIdentifierError: string = "", responseError = "";
    let showPassword: boolean;

    async function handleLogin() {

        passwordError = "";
        userIdentifierError = "";
        responseError = "";

        const passErr = validatePassword(password);

        const uidErr = validateUserIdentifier(userIdentifier);

        if (passErr) {
            passwordError = passErr.message;
        }
        if (uidErr) {
            userIdentifierError = uidErr.message;
        }

        if (passErr || uidErr) {
            return;
        }

        const request: UserSignInRequest = {
            uid: userIdentifier,
            password: password
        };

        const response = await api.AuthService.SignIn(request);

        if (response.error) {
            responseError = response.error.message;
            return;
        }

        //api.AccessToken = <string>response.response?.accessToken
        goto("/");
    }

    export function handleShowHidePassword() {
        showPassword = !showPassword;
    }

</script>

<div id="page">

    <Input label="username / email"
           type="text"
           placeholder="username / email"
           bind:value={userIdentifier}
           error={!!userIdentifierError}
    />
    <Input label="password"
           type={showPassword ? "text": "password"}
           placeholder="password"
           bind:value={password}
           error={!!passwordError}
    />

    <Checkbox className="m-auto" onclick={handleShowHidePassword}>Show password</Checkbox>

    <ErrorTextBox className="m-auto" bind:value={responseError}/>

    <Button className="m-auto" onClick={handleLogin}>Login</Button>
</div>

<style>
    #page {
        display: flex;
        justify-content: center;
        flex-direction: column;
    }
</style>
