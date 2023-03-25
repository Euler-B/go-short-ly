<script>
    import Card from "./Card.svelte";
    import Modal from "./Modal.svelte";
    import { Modals, closeModal, openModal } from "svelte-modals";

    export let goly 
    let ShowCard = true

    async function update(data) {
        const json = {
        redirect: data.redirect,
        goly: data.goly,
        random: data.random,
        id: goly.id
        }

        await fetch("http://localhost:3000/goly", {
            method: "PATCH",
            headers: { "Content-Type": "application/json"},
            body: JSON.stringify(json)
        }).then(res => {
            console.log(res)
        })
    }

    function handleOpen(goly) {
        openModal(Modal, {
            title: "Update Goly link",
            send: update,
            goly: goly.goly,
            redirect: goly.redirect,
            random: goly.random
        })
    }

    async function deleteGoly() {
        if (confirm("Are you sure you wish to delelete this Goly link (" + goly.goly + ") ? ")) {
            await fetch("http://localhost:3000/goly/" + goly.id, {
                method: "DELETE"
            }).then(response => {
                ShowCard = false
                console.log(response)
            })
        }
    }
    
</script>

{#if ShowCard }
<Card>
    <p>Goly: http:localhost:3000/r/{goly.goly}</p>
    <p>Redirect: {goly.redirect}</p>
    <p>Clicked: {goly.clicked}</p>
    <button class="update" on:click={ handleOpen(goly) }>Update</button>
    <button class="delete" on:click={ deleteGoly }>Delete</button>
</Card>
{/if}

<Modals>
    <div 
        slot="backdrop"
        class="backdrop"
        on:click={closeModal}
    />
</Modals>

<style>
    button {
        color: white;
        font-weight: bolder;
        border: none;
        padding: .75rem;
        border-radius:4px;
    }
    .update {
        background-color: yellowgreen;
    }
    .delete {
        background-color: red;
    }
</style>