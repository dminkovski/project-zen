import {
  createStyles,
  Group,
  Paper,
  Text,
  ThemeIcon,
  rem,
} from "@mantine/core";
import { IconMail } from "@tabler/icons-react";

const useStyles = createStyles((theme) => ({
  card: {
    position: "relative",
    cursor: "pointer",
    overflow: "hidden",
    transition: "transform 150ms ease, box-shadow 100ms ease",
    padding: theme.spacing.xl,
    marginBottom: "50px",
    paddingLeft: `calc(${theme.spacing.xl} * 2)`,

    "&:hover": {
      boxShadow: theme.shadows.md,
      transform: "scale(1.02)",
    },

    "&::before": {
      content: '""',
      position: "absolute",
      top: 0,
      bottom: 0,
      left: 0,
      width: rem(6),
      backgroundImage: theme.fn.linearGradient(
        0,
        theme.colors.green[5],
        theme.colors.gray[2]
      ),
    },
  },
}));

interface EmailProps {
  title: string;
  body: string;
  date: string;
  from: string;
}

export function Email({ title, body, date, from }: EmailProps) {
  const { classes } = useStyles();
  return (
    <Paper withBorder radius="md" className={classes.card}>
      <Group position="center" spacing="sm">
        <Text
          size="l"
          weight={500}
          mt="md"
          variant="gradient"
          gradient={{ from: "black", to: "gray", deg: 45 }}
        >
          {from}
        </Text>
      </Group>

      <Text color="dimmed" size="xs" transform="uppercase" weight={700} mt="md">
        {date}
      </Text>
      <Text size="xl" weight={500} mt="md">
        {title}
      </Text>
      <Text
        size="sm"
        mt="sm"
        color="dimmed"
        align="justify"
        dangerouslySetInnerHTML={{ __html: body }}
      ></Text>
    </Paper>
  );
}
