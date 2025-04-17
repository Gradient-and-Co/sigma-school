import axios from "axios";

export async function apiRequest(url: string, method: string = 'get', payload?: object, token?: string) {
    let headers = {}
    if (token) {
        headers = {'Authorization': token}
    }

    let data, error
    await axios({
        method: method,
        url: "http://95.165.5.237:8080/api/v1" + url,
        withCredentials: true,
        data: payload,
        headers: headers,
    })
    .then(resp => {
        data = resp.data;
    })
    .catch(err => {
        error = err.response.data
    });

    return {data, error}
}

export async function uploadFile(e: Event) {
    const input = e.target as HTMLInputElement;
    let data, error
    if (input.files && input.files.length > 0) {
        const formData = new FormData();
        formData.append("file", input.files[0]);

        await axios({
            method: 'post',
            url: 'http://95.165.5.237:8080/api/v1/upload',
            withCredentials: true,
            data: formData,
        })
            .then(resp => {
                data = resp.data;
            })
            .catch(err => {
                error = err.response.data
            });
    } else {
        error = 'file is not available';
    }

    return {data, error}
}
