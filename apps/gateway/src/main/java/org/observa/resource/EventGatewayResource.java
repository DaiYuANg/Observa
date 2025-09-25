package org.observa.resource;

import io.smallrye.mutiny.Uni;
import jakarta.ws.rs.POST;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.core.Response;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.eclipse.microprofile.openapi.annotations.tags.Tag;
import org.observa.model.EventDTO;
import org.observa.service.EventService;

@Path("/api/v1/events")
//@Consumes(MediaType.APPLICATION_JSON)
//@Produces(MediaType.APPLICATION_JSON)
@RequiredArgsConstructor
@Slf4j
@Tag(name = "event gateway")
public class EventGatewayResource {

  private final EventService eventService;

  @POST
  @Path("/")
  public Uni<Response> receiveEvent(EventDTO event) {
    log.atInfo().log("Received event: " + event);
    eventService.sendEvent(event);
    return Uni.createFrom().item(Response.status(Response.Status.ACCEPTED).build());
  }
}