<nav>
    <a href="." id="logo">
        <img id="logo" src="logo.png" alt="Encedeus logo">
    </a>


    <input id="searchBar" class="navbarComponent" type="search" placeholder="Search plugins">

    <div>
        <a href="./documentation">
            <button id="docs" class="navbarComponent">Documentation</button>
        </a>
    </div>

    <div id="authButtons">
        <button id="signUp" class="navbarComponent">Sign Up</button>

        <button id="signIn" class="navbarComponent">Sign In</button>
    </div>

</nav>

<style>
    nav {
        top: 0;
        left: 0;
        margin-top: 0;
        padding-bottom: 5px;

        position: relative;

        background-color: gray;

        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: center;
        align-items: center;

        gap: 5px 20px;
    }

    #logo,
    #searchBar,
    #authButtons {
        max-height: 100px;
        display: flex;
        flex-direction: row;
        justify-content: center;
        align-items: center;
        gap: inherit;
    }

    #searchBar {
        min-width: 20%;
        width: 50%;
        border: 0;
        padding: 0;
    }

    button {
        max-height: 100px;
        display: flex;
        flex-direction: row;
        justify-content: center;
        align-items: center;

        border: 0;
        border-bottom: 50px;
    }

    .navbarComponent {
        height: 45px;
    }

</style>