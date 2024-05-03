$('#login').on('submit', fazerLogin);

function fazerLogin (evento){
    evento.preventDefault(); // prefine  comportamento default do formulario que recarrega a página

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(),
            senha: $('#senha').val()
        }
    }).done(function(){ // Ele sabe se funcionou com base no valor do status code
        window.location = "/home"
    }).fail(function(erro){
        console.log(erro)
        alert("Usuário ou senha inválidos");
    })

}