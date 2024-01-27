<script lang="ts">
    import Input from "$lib/components/Input.svelte";
    import Button from "$lib/components/Button.svelte";
    import {userDataStore} from "$lib/stores/userDataStore";
    import {getApi} from "$lib/api/api";
    import {UserUpdateRequest} from "@encedeus/registry-js-api";
    import ErrorTextBox from "$lib/components/ErrorTextBox.svelte";
    import {goto} from "$app/navigation";

    let newUsername: string = "",
        currentEmail: string = $userDataStore.email,
        newEmail: string = "",
        newPassword: string = "",
        passwordConfirmation: string = "",
        errorMessage: string = "",
        isLoading: boolean = false,
        currentName: string = $userDataStore.name;

    const api = getApi();


    async function handleSaveData() {
        errorMessage = "";

        const req = {
            name: newUsername,
            email: newEmail,
            password: newPassword
        } as UserUpdateRequest;

        isLoading = true;

        const resp = await api.UsersService.UpdateUser(req);

        isLoading = false;

        if (resp.error) {
            errorMessage = resp.error.message;
            return;
        }
    }

    function handleDiscardData() {
        newUsername = "";
        newEmail = "";
        newPassword = "";
        passwordConfirmation = "";
    }

    async function handleDeleteUser() {
        const resp = await api.UsersService.DeleteSelf()

        if (resp.error) {
            errorMessage = resp.error.message;
            return;
        }

        goto("/")
    }

</script>

<div id="page">
    <div>
        <Input helperText="new username" bind:value={newUsername} placeholder={currentName}/>
    </div>

    <div>
        <Input helperText="new email" bind:value={newEmail} placeholder={currentEmail}/>
    </div>

    <div>
        <Input helperText="new password" bind:value={newPassword}/>
        <Input helperText="confirm password" bind:value={passwordConfirmation}/>
    </div>

    <Button onclick={handleSaveData}>Save</Button>
    <Button onclick={handleDiscardData}>Discard changes</Button>
    <Button onclick={handleDeleteUser}>Delete account</Button>

    <!-- state to be used for a loading indicator-->
    <p class:invisible={!isLoading}>loading...</p>

    <ErrorTextBox bind:value={errorMessage}/>
</div>

<style>
    #page {
        text-align: center;
    }

    .invisible {
        display: none;
    }
</style>