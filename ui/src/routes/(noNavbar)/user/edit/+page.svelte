<script lang="ts">
    import Input from "$lib/components/generic/Input.svelte";
    import Button from "$lib/components/generic/Button.svelte";
    import {userDataStore} from "$lib/stores/userDataStore";
    import {getApi} from "$lib/api/api";
    import {UserUpdateRequest} from "@encedeus/registry-js-api";
    import ErrorTextBox from "$lib/components/ErrorTextBox.svelte";
    import {goto} from "$app/navigation";
    import Checkbox from "$lib/components/Checkbox.svelte";

    let newUsername: string = "",
        currentEmail: string = $userDataStore.email,
        newEmail: string = "",
        newPassword: string = "",
        passwordConfirmation: string = "",
        errorMessage: string = "",
        isLoading: boolean = false,
        currentName: string = $userDataStore.name,
        showPassword: boolean

    const api = getApi();


    async function handleSaveData() {
        if (passwordConfirmation != newPassword) {
            errorMessage = "make sure the password is entered correctly"
            return
        }

        const req = {
            name: newUsername,
            email: newEmail,
            password: newPassword
        } as UserUpdateRequest;
       if (!req.name && !req.password && !req.email) {
           errorMessage = "enter new parameters"
           return
       }

        isLoading = true;

        const resp = await api.UsersService.UpdateUser(req);

        isLoading = false;

        if (resp.error) {
            errorMessage = resp.error.message;
            return;
        }

        errorMessage = "";
    }

    function handleDiscardData() {
        newUsername = "";
        newEmail = "";
        newPassword = "";
        passwordConfirmation = "";
    }

    async function handleDeleteUser() {
        const resp = await api.UsersService.DeleteSelf();

        console.log(resp);

        if (resp.error) {
            errorMessage = resp.error.message;
            return;
        }

        goto("/");
    }

    function handleShowHidePassword() {
        showPassword = !showPassword;
    }

</script>

<div id="page">
    <div>
        <Input label="new username" bind:value={newUsername} placeholder={currentName}/>
    </div>

    <div>
        <Input label="new email" bind:value={newEmail} placeholder={currentEmail}/>
    </div>

    <div>
        <Input type={showPassword ? "text": "password"} placeholder="new password" label="new password" bind:value={newPassword}/>
        <Input type={showPassword ? "text": "password"} placeholder="confirm new password" label="confirm new password" bind:value={passwordConfirmation}/>
    </div>

    <ErrorTextBox bind:value={errorMessage}/>

    <Checkbox onclick={handleShowHidePassword}>Show password</Checkbox>

    <Button className="mt-2.5" onClick={handleSaveData}>Save</Button>
    <Button className="mt-2.5" onClick={handleDiscardData}>Discard changes</Button>
    <br>
    <Button className="mt-2.5 bg-red-500 hover:bg-red-600 active:bg-red-700" onClick={handleDeleteUser}>Delete account</Button>
</div>

<style>
    #page {
        text-align: center;
    }

    .invisible {
        display: none;
    }
</style>