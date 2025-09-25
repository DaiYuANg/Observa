package org.observa;

import jakarta.ws.rs.GET;
import jakarta.ws.rs.Path;

@Path("/api/v1/auth")
public class AuthResource {

    @GET
    public String hello() {
        return "Hello from Auth Resource";
    }
}