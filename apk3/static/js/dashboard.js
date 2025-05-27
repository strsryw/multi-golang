$(document).ready(function(){
    // alert("js berjalan !!");
    getData();
});

function resetInput(){
   
    $('#inpId').val('');
    $('#inpNama').val('');
    $('#inpStatus').val('');
    $('#inpAktif').val('');
    $('#inpCons').val('');
}

function getData(){
    $.ajax({
        url:'/apk3',
        type:'post',
        data:{action:'getData'},
        success:function(response){
            // console.log(response);
            $('#tblGetData').children('tbody:first').html(response);
        }
    })
}

function edit(id){
    $('#inpId').val(id);
    $('#inpNama').val($('#txtNama'+id).text());
    $('#inpStatus').val($('#txtStatus'+id).text());
    $('#inpAktif').val($('#txtAktif'+id).text());
    $('#inpCons').val($('#txtCons'+id).text());
}

function hapus(id){
    if(confirm("Yakin ingin menghapus data ini ?")){
        $.ajax({
            url:'/apk3',
            type:'post',
            data:{id, action:"hapus"},
            success:function(response){
                console.log(response);
                response = response.split("~!~");
                alert(response[1]);
                if(response[0] == "sukses"){
                    getData();
                }else{
                   
                }
               
            }
        })
    }
}


function simpan(){
    let inpId = $('#inpId').val();
    let inpNama = $('#inpNama').val();
    let inpStatus = $('#inpStatus').val();
    let inpAktif = $('#inpAktif').val();
    let inpCons = $('#inpCons').val();

    $.ajax({
        url:'/apk3',
        type:'post',
        data:{inpId, inpNama, inpStatus, inpAktif, inpCons, action:"simpan"},
        success:function(response){
            console.log(response);
            response = response.split("~!~");
            if(response[0] == "sukses"){
                alert(response[1]);
                $('#inpId').val('');
                $('#inpNama').val('');
                $('#inpStatus').val('');
                $('#inpAktif').val('');
                $('#inpCons').val('');
                getData();
            }else{
                alert(response[1]);
            }
           
        }
    });
}

