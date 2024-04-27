$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario (evento){
    evento.preventDefault();

    console.log("hello")
    console.log($('#senha').val())
    console.log($('#confirmar-senha').val())


    if ($('#senha').val() != $('#confirmar-senha').val()) {
        alert("As senhas não coincidem!")
        return
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val()
        }
    }).done(function(){ // Ele sabe se funcionou com base no valor do status code
        alert("Usuário cadastrado com sucesso");
    }).fail(function(erro){
        console.log(erro)
        alert("Erro ao cadastrar o usuário");
    })


}