package br.com.rocketseat.hiokdev.planner_java.config.mail;

import jakarta.validation.constraints.NotNull;
import lombok.Getter;
import lombok.Setter;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Component;
import org.springframework.validation.annotation.Validated;

@Getter
@Setter
@Validated
@Component
@ConfigurationProperties("planner.mail")
public class MailProperties {

    @NotNull
    private String sender;

}